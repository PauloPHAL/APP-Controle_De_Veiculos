import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { RoutesService } from './routes.service';
import { CreateRouteDto } from './dto/create-route.dto';
import { UpdateRouteDto } from './dto/update-route.dto';
import { RoutesDriverService } from './routes-driver/routes-driver.service';

@Controller('routes')
export class RoutesController {
  constructor(
    private readonly routesService: RoutesService,
    private routesDriverService: RoutesDriverService,
  ) { }

  @Post()
  create(@Body() createRouteDto: CreateRouteDto) {
    return this.routesService.create(createRouteDto);
  }

  @Get()
  findAll() {
    return this.routesService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.routesService.findOne(id);
  }

  @Post(':id/process-route')
  processRoute(@Param('id') id: string, @Body() payload: { lat: number, lng: number }) {
    return this.routesDriverService.processRoute({
      route_id: id,
      lat: payload.lat,
      lng: payload.lng,
    });
  }

  @Post(':id/start')
  startRoute(@Param('id') id: string) {
    return this.routesService.startRoute(id);
  }
}
