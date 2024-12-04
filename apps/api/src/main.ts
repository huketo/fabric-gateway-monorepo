import { Logger, ClassSerializerInterceptor } from '@nestjs/common';
import { NestFactory, Reflector } from '@nestjs/core';
import { AppModule } from './app.module';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalInterceptors(new ClassSerializerInterceptor(app.get(Reflector)));

  const globalPrefix = 'api';
  app.setGlobalPrefix(globalPrefix);
  const port = process.env.PORT || 3000;

  const documentConfig = new DocumentBuilder()
    .setTitle('Hyperledger Fabric Asset Management API')
    .setDescription('API for managing assets on Hyperledger Fabric')
    .setVersion('1.0.0')
    .addServer(
      `http://localhost:${port}/${globalPrefix}`,
      'Local development server'
    )
    .setOpenAPIVersion('3.0.0')
    .build();

  const document = SwaggerModule.createDocument(app, documentConfig);
  SwaggerModule.setup('api', app, document, {
    explorer: true,
    swaggerUiEnabled: true,
    jsonDocumentUrl: '/api/swagger.json',
  });

  await app.listen(port);

  Logger.log(
    `🚀 Application is running on: http://localhost:${port}/${globalPrefix}`
  );
}

bootstrap();
