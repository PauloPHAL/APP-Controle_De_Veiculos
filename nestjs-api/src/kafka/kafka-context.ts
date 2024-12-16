import * as Kafka from '@confluentinc/kafka-javascript';

export class KafkaContext {
    constructor(
        readonly message: Kafka.KafkaJS.Message,
        readonly messageValue: any,
        readonly topic: string,
        readonly partition: number,
        readonly consumer: Kafka.KafkaJS.Consumer,
    ) { }
}