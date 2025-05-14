import Container from "@/components/Container";
import { Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <>
      <Container>
        <Outlet />
      </Container>
    </>
  ),
});
