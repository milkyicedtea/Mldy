import {homeStyles} from "../styles/homeStyles.ts"
import {Anchor, AppShell, Container, Text} from "@mantine/core"
import {useNavigate} from "@tanstack/react-router"

export function Header() {
  const navigate = useNavigate();

  return (
    <>
      <AppShell.Header p={4} style={{ display: "flex" }}>
        <Container style={{...homeStyles.titleContainer}}>
          <Anchor
            href={"/"}
            style={{...homeStyles.anchor}}
            onClick={(e) => {
              e.preventDefault()
              void navigate({to: "/"})
            }}
          >
            <Text gradient={{from: 'grape', to: 'blue', deg: 90}} variant={'gradient'}
              fw={400} style={{...homeStyles.title}}
            >
              Mldy
            </Text>
          </Anchor>
        </Container>
      </AppShell.Header>

    </>
  )
}