import {createRootRoute, createRoute, createRouter} from "@tanstack/react-router"
import {Home} from "./pages/Home.tsx"
import {Disclaimer} from "./pages/Disclaimer.tsx"

const rootRoute = createRootRoute()

const homeRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Home,
})

const disclaimerRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/disclaimer",
  component: Disclaimer,
})

const routeTree = rootRoute.addChildren([homeRoute, disclaimerRoute])
const router = createRouter({routeTree})

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

export { router }
