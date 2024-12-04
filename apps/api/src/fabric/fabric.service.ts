import { Injectable, Logger, OnApplicationBootstrap } from '@nestjs/common';
import * as grpc from '@grpc/grpc-js';
import {
  connect,
  Gateway,
  hash,
  Identity,
  Signer,
  signers,
} from '@hyperledger/fabric-gateway';
import * as crypto from 'crypto';
import * as path from 'path';
import { promises as fs } from 'fs';

@Injectable()
export class FabricService implements OnApplicationBootstrap {
  public geteway: Gateway;

  async onApplicationBootstrap() {
    const cryptoPath = path.resolve(__dirname, 'configs', 'org1', 'user1');
    const keyDirectoryPath = path.resolve(cryptoPath, 'keystore');
    const certDirectoryPath = path.resolve(cryptoPath, 'signcerts');
    const tlsCertPath = await this.getFirstDirFileName(
      path.resolve(cryptoPath, 'cacerts')
    );
    const peerEndpoint = 'localhost:7051';
    const peerHostAlias = 'peer0.org1.example.com';

    const client = await this.newGrpcConnection(
      tlsCertPath,
      peerEndpoint,
      peerHostAlias
    );
    const org1Identity = await this.newIdentity('Org1MSP', certDirectoryPath);
    const org1Signer = await this.newSigner(keyDirectoryPath);

    this.geteway = connect({
      client,
      identity: org1Identity,
      signer: org1Signer,
      hash: hash.sha256,
      // Default timeouts for different gRPC calls
      evaluateOptions: () => {
        return { deadline: Date.now() + 5000 }; // 5 seconds
      },
      endorseOptions: () => {
        return { deadline: Date.now() + 15000 }; // 15 seconds
      },
      submitOptions: () => {
        return { deadline: Date.now() + 5000 }; // 5 seconds
      },
      commitStatusOptions: () => {
        return { deadline: Date.now() + 60000 }; // 1 minute
      },
    });

    Logger.log('Fabric service started');
  }

  private async newGrpcConnection(
    tlsCertPath: string,
    peerEndpoint: string,
    peerHostAlias: string
  ): Promise<grpc.Client> {
    const tlsRootCert = await fs.readFile(tlsCertPath);
    const tlsCredentials = grpc.credentials.createSsl(tlsRootCert);
    return new grpc.Client(peerEndpoint, tlsCredentials, {
      'grpc.ssl_target_name_override': peerHostAlias,
    });
  }

  private async getFirstDirFileName(dirPath: string): Promise<string> {
    const files = await fs.readdir(dirPath);
    const file = files[0];
    if (!file) {
      throw new Error(`No files found in ${dirPath}`);
    }
    return path.join(dirPath, file);
  }

  private async newIdentity(
    mspId: string,
    certDirectoryPath: string
  ): Promise<Identity> {
    const certPath = await this.getFirstDirFileName(certDirectoryPath);
    const credentials = await fs.readFile(certPath);
    return { mspId, credentials };
  }

  private async newSigner(keyDirectoryPath: string): Promise<Signer> {
    const keyPath = await this.getFirstDirFileName(keyDirectoryPath);
    const privateKeyPem = await fs.readFile(keyPath);
    const privateKey = crypto.createPrivateKey(privateKeyPem);
    return signers.newPrivateKeySigner(privateKey);
  }
}
