import {AppShell, Container, Text, Title} from "@mantine/core"
import {Footer} from "../components/Footer.tsx"
import {Header} from "../components/Header.tsx"
import {homeStyles} from "../styles/homeStyles.ts"

export function Disclaimer() {
  return (
    <>
      <Header/>
      <AppShell.Main>
        <Container style={{...homeStyles.contentContainer, marginTop: "6.25%"}}>
          <Title>Disclaimer & Terms of Use</Title>
          <Text>
            <strong>Mldy/GoMldy</strong> is an open-source, non-commercial tool designed for personal use only.
            It does not use any official services' API and is not affiliated with or endorsed by any platform it may support.
            Users are solely responsible for ensuring that their use of this tool complies with applicable laws,
            terms of service, and copyright regulations.
          </Text>
          <Text>
            <strong>Privacy Policy</strong> <br/>
            Mldy/GoMldy does not collect or permanently store user data. IP logs are used only for rate-limiting purposes and
            are automatically purged after their cooldown period. No tracking, analytics, or third-party data sharing occurs.
            The collection of non-permanently stored data only happens on the web version (Mldy). The local application
            (GoMldy) doesn't store any user data at all, besides the content downloaded
          </Text>
          <Text>
            <strong>Terms of Use</strong> <br/>
            - This tool is intended only for legal, personal use.
            - The developers and maintainers of Mldy/GoMldy are not responsible for any misuse.
            - Automated usage (bots, mass downloading, excessive requests) is **not allowed**.
          </Text>
          <Text>
            <strong>Legal Disclaimer</strong> <br/>
            - Mldy/GoMldy is not responsible for how users interact with external services.
            - Users must ensure that they have the right to download content.
            - Mldy/GoMldy should not be used for copyrighted material without proper authorization.
          </Text>
          <Text>
            This project is provided <strong>as-is</strong>, without warranty of any kind.
          </Text>
        </Container>
      </AppShell.Main>
      <Footer/>
    </>
  )
}