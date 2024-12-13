import { Controller, Get, Query } from '@nestjs/common';
import { DirectionsService } from './directions.service';

@Controller('directions')
export class DirectionsController {

    constructor(private directionsService: DirectionsService) {}

    @Get()
    getDirections(@Query('originID') originID: string, @Query('destinationID') destinationID: string) {
        return this.directionsService.getDirections(originID, destinationID);
    }
}
