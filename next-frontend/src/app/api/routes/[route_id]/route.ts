//routes
import { NextRequest, NextResponse } from "next/server";

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ route_id: string }> }
) {
  const { route_id } = await params;
  const response = await fetch(`${process.env.NEST_API_URL}/routes/${route_id}`, {
    cache: "force-cache",
    next: {
      tags: [`routes-${route_id}`, "routes"],
    },
  });
  const data = await response.json();
  return NextResponse.json(data);
}
