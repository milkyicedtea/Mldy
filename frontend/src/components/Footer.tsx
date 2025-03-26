import {Anchor, AppShell, Container, Text} from "@mantine/core"
import {Link} from "@tanstack/react-router"
import {homeStyles} from "../styles/homeStyles.ts"
import {openExternalLink} from "@local/hooks/openExternalLink.ts";

export function Footer() {

  return (
    <>
      <AppShell.Footer zIndex={0} p={"sm"}>
        <Container style={{...homeStyles.footerContainer}}>
          <Anchor
            href="https://051205.xyz"
            style={{...homeStyles.anchor, color: 'white'}}
            onClick={(e) => {
              e.preventDefault()
              openExternalLink("https://051205.xyz")
            }}
          >
            <Text variant={"text"}>© 2025 051205.xyz</Text>
          </Anchor>

          <Link style={{color: 'aliceblue'}} to={"/disclaimer"}>Disclaimer</Link>
        </Container>
      </AppShell.Footer>
    </>
  )
}