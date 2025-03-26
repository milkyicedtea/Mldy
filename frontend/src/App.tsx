import {router} from "./router.tsx";
import {RouterProvider} from "@tanstack/react-router";
import {AppShell, MantineProvider} from "@mantine/core";
import {theme} from "./theme.ts";
import '@mantine/core/styles.css'

function App() {
  return (
    <MantineProvider theme={theme} defaultColorScheme={'auto'}>
      <AppShell padding={'sm'}>
        <RouterProvider router={router}/>
      </AppShell>
    </MantineProvider>
  )
}

export default App
