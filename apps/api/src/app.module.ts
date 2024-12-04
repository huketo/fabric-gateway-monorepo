import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { FabricModule } from './fabric/fabric.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    FabricModule,
  ],
})
export class AppModule {}
