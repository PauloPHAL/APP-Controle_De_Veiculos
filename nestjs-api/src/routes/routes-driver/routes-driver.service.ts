import { Get, Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';
import { RoutesDriverGateway } from './routes-driver.gateway';

@Injectable()
export class RoutesDriverService {

    constructor(private prismaService: PrismaService, private routesGateway: RoutesDriverGateway) { }

    async processRoute(dto: { route_id: string; lat: number; lng: number; }) {
        const routeDriver = await this.prismaService.routeDriver.upsert({
            include: {
                route: true,
            },
            where: { route_id: dto.route_id },
            create: {
                route_id: dto.route_id,
                points: {
                    set: {
                        location: {
                            lat: dto.lat,
                            lng: dto.lng
                        }
                    }
                },
            },
            update: {
                points: {
                    push: {
                        location: {
                            lat: dto.lat,
                            lng: dto.lng
                        }
                    }
                }
            }
        });

        this.routesGateway.emitNewPoints({
            route_id: dto.route_id,
            lat: dto.lat,
            lng: dto.lng,
        });

        return routeDriver;
    };
}
