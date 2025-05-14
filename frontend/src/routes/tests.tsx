import { createFileRoute } from "@tanstack/react-router";
import TestView from "@/views/tests";

export const Route = createFileRoute("/tests")({
  component: RouteComponent,
});

function RouteComponent() {
  return <TestView />;
}
