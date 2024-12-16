import { Controller, Logger } from '@nestjs/common';
import { MessagePattern } from '@nestjs/microservices';
import { KafkaContext } from '../../kafka/kafka-context';
import { HttpService } from '@nestjs/axios';

@Controller()
export class RoutesDriverConsumer {
  private logger = new Logger(RoutesDriverConsumer.name);

  constructor(private httpService: HttpService) {}

  @MessagePattern('simulation')
  async driverMoved(payload: KafkaContext) {
    this.logger.log(
      `Receiving message from topic ${payload.topic}`,
      payload.messageValue,
    );
    
    const { route_id, lat, lng } = payload.messageValue;
    await this.httpService.axiosRef.post(
      `http://localhost:3000/routes/${route_id}/process-route`,
      {
        lat,
        lng,
      },
    );
  }
}