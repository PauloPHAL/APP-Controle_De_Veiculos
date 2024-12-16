"use client";

import { useEffect, useRef } from "react";
import { socket } from "../../utils/socket-io";
import { useMap } from "@/hooks/useMaps";

export type MapDriverProps = {
    routeIdElementId: string;
};

export function MapDriver(props: MapDriverProps) {
    const { routeIdElementId } = props;
    const mapContainerRef = useRef<HTMLDivElement>(null);
    const map = useMap(mapContainerRef);

    useEffect(() => {
        if (!map || !routeIdElementId) {
            return;
        }

        const selectElement = document.querySelector(
            `#${routeIdElementId}`
        )!;

        socket.connect();

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const handler = async (event: any) => {
            console.log('aaaaaa');
            socket.offAny();
            const routeId = event.target!.value;
            socket.on(
                `server:new-points/${routeId}:list`,
                async (data: { route_id: string; lat: number; lng: number }) => {
                    if (!map.hasRoute(data.route_id)) {
                        const response = await fetch(
                            `${process.env.NEXT_PUBLIC_NEXT_API_URL}/routes/${data.route_id}`
                        );
                        const route = await response.json();
                        map.addRouteWithIcons({
                            routeId: data.route_id,
                            startMarkerOptions: {
                                position: route.directions.routes[0].legs[0].start_location,
                            },
                            endMarkerOptions: {
                                position: route.directions.routes[0].legs[0].end_location,
                            },
                            carMarkerOptions: {
                                position: route.directions.routes[0].legs[0].start_location,
                            },
                        });
                    }
                    map.moveCar(data.route_id, { lat: data.lat, lng: data.lng });
                }
            );
        }

        selectElement.addEventListener("change", handler);

        return () => {
            selectElement.removeEventListener("change", handler);
            socket.disconnect();
        };
    }, [routeIdElementId, map]);

    return (
        <div
            className="w-2/3 h-full"
            ref={mapContainerRef}
        ></div>
    );
}