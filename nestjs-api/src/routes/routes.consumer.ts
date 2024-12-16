import { Controller, Logger } from "@nestjs/common";
import { MessagePattern } from "@nestjs/microservices";
import { KafkaContext } from "src/kafka/kafka-context";
import { RoutesService } from "./routes.service";

@Controller()
export class RoutesConsumer {
    private logger = new Logger(RoutesConsumer.name);

    constructor(private routesServices: RoutesService) { }

    @MessagePattern('freight')
    async updateFreight(payload: KafkaContext) {
        this.logger.log(
            `Receiving message from topic ${payload.topic}`,
            payload.messageValue,
        );
        const { route_id, amount } = payload.messageValue;

        await this.routesServices.updateFreight(route_id, { freight: amount });
    }
}