import { Inject, Module, OnModuleInit } from '@nestjs/common';
import * as Kafka from '@confluentinc/kafka-javascript'
import { ConfigService } from '@nestjs/config';

@Module({
    providers: [
        {
            provide: 'KAFKA_PRODUCER',
            useFactory: (config: ConfigService) => {
                return new Kafka.KafkaJS.Kafka({
                    'bootstrap.servers': config.get('KAFKA_BROKER'),
                }).producer();
            },
            inject: [ConfigService],
        }
    ],
    exports: ['KAFKA_PRODUCER'],
})
export class KafkaModule implements OnModuleInit{
    constructor(
        @Inject('KAFKA_PRODUCER') private kafkaProducer: Kafka.KafkaJS.Producer,
      ) {}

    async onModuleInit() {
        await this.kafkaProducer.connect();
    }
}
