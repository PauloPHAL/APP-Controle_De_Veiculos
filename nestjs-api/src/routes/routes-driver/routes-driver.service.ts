import { Get, Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class RoutesDriverService {

    constructor(private prismaService: PrismaService) {}

    
    processRoute(dto: {router_id: string; lat: number; lng: number;}) {
        return this.prismaService.routeDriver.upsert({
            include:{
                router: true,
            },
            where: {router_id: dto.router_id},
            create: {
                router_id: dto.router_id,
                points: {
                    set:{
                        location:{
                            lat: dto.lat,
                            lng: dto.lng
                        }
                    }
                },
            },
            update: {
                points:{
                    push:{
                        location:{
                            lat: dto.lat,
                            lng: dto.lng
                        }
                    }
                }
            }
        });
    };
}
