import {Anchor, Container, Grid, Image, Text, Title} from "@mantine/core"
import {homeStyles} from "../styles/homeStyles.ts"
import {openExternalLink} from "@local/hooks/openExternalLink.ts";
import facebookSvg from "../assets/facebook.svg"
import instagramSvg from "../assets/instagram.svg"
import redditSvg from "../assets/reddit.svg"
import soundcloudSvg from "../assets/soundcloud.svg"
import XSvg from "../assets/X.svg"
import youtubeSvg from "../assets/youtube.svg"

export function SupportedPlatforms() {
  return (
    <>
      <Container style={{ display: "flex", flexDirection: "column", marginTop: "5%", padding: 0, marginLeft: 0 }}>
        <Title>Supported platforms</Title>
        <Grid style={{ padding: 0, margin: 0, marginBottom: "1.5%" }}>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://facebook.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://facebook.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={facebookSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://instgram.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://instagram.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={instagramSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://reddit.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://reddit.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={redditSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://soundcloud.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://soundcloud.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={soundcloudSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://x.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://x.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={XSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
          <Grid.Col span={2}>
            <Container style={{...homeStyles.iconContainer}}>
              <Anchor href={"https://youtube.com"}
                onClick={(e) => {
                  e.preventDefault()
                  openExternalLink("https://youtube.com")
                }}
              >
                <Image style={{...homeStyles.platformIcon}} src={youtubeSvg} />
              </Anchor>
            </Container>
          </Grid.Col>
        </Grid>
        <Text variant={'text'} style={{}}>(And More!)</Text>
      </Container>
    </>
  )
}