import { Module } from '@nestjs/common';
import { RoutesService } from './routes.service';
import { RoutesController } from './routes.controller';
import { MapsModule } from 'src/maps/maps.module';
import { RoutesDriverService } from './routes-driver/routes-driver.service';
import { RoutesDriverGateway } from './routes-driver/routes-driver.gateway';
import { KafkaModule } from 'src/kafka/kafka.module';
import { RoutesConsumer } from './routes.consumer';
import { HttpModule } from '@nestjs/axios';
import { RoutesDriverConsumer } from './routes-driver/routes-driver.consumer';

@Module({
  imports: [MapsModule, KafkaModule, HttpModule],
  controllers: [RoutesController, RoutesConsumer, RoutesDriverConsumer],
  providers: [RoutesService, RoutesDriverService, RoutesDriverGateway],
})
export class RoutesModule {}
