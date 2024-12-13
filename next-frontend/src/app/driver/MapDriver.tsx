"use client";

import { useMap } from "@/hooks/useMaps";
import { socket } from "@/utils/socket-io";
import { useEffect, useRef } from "react";

export type MapDriverProps = {
    route_id: string | null;
    start_location: {
        lat: number;
        lng: number;
    } | null;
    end_location: {
        lat: number;
        lng: number;
    } | null;
};

export function MapDriver(props: MapDriverProps) {
    const { route_id, start_location, end_location } = props;
    const mapContainerRef = useRef<HTMLDivElement>(null);
    const map = useMap(mapContainerRef);

    useEffect(() => {
        if (!map || !route_id || !start_location || !end_location) {
            return;
        }

        // eslint-disable-next-line @typescript-eslint/no-unused-expressions
        socket.disconnected ? socket.connect() : socket.offAny();

        socket.on("connect", () => {
            console.log("connected");
            socket.emit(`client:new-points`, { route_id });
        });

        socket.on(
            `server:new-points/${route_id}:list`,
            (data: { route_id: string; lat: number; lng: number }) => {
        
                if (!map.hasRoute(data.route_id)) {
                    map.addRouteWithIcons({
                        routeId: data.route_id,
                        startMarkerOptions: {
                            position: start_location,
                        },
                        endMarkerOptions: {
                            position: end_location,
                        },
                        carMarkerOptions: {
                            position: start_location,
                        },
                    });
                }
                map.moveCar(data.route_id, { lat: data.lat, lng: data.lng });
            }
        );

        return () => {
            socket.disconnect();
        };

    }, [route_id, start_location, end_location, map]);

    return (
        <div
            className="w-2/3 h-full"
            ref={mapContainerRef}
        ></div>
    );
}