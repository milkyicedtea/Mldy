import {AppShell, Button, Container, Loader, Text, TextInput} from "@mantine/core"
import {useState} from "react"
import {Footer} from "@local/components/Footer.tsx"
import {Header} from "@local/components/Header.tsx"
import {homeStyles} from "@local/styles/homeStyles.ts"
import {SupportedPlatforms} from "../components/SupportedPlatforms.tsx"
import {Download} from "@wailsjs/go/main/App";
import {main} from "@wailsjs/go/models.ts";
import VideoRequest = main.VideoRequest;

export function Home() {
  const [url, setUrl] = useState<string | null>(null)
  const [downloading, setDownloading] = useState(false)
  const [downloadedPath, setDownloadedPath] = useState<string|null>(null)

  async function download(url: string) {
    if (downloading) return
    try {
      setDownloadedPath(null)
      setDownloading(true)


      setDownloadedPath(await Download(VideoRequest.createFrom({url})))
      setDownloading(false)
    } catch (error: any) {
      setDownloading(false)
      // console.log(error.response.data)
      alert(`${await error}`)
      console.error(error)
    }
  }

  return (
    <>
      <Header/>
      <AppShell.Main>
        <Container style={{...homeStyles.contentContainer}}>
          <SupportedPlatforms/>
          {/*<Container style={{...homeStyles.infoGroupContainer}}>*/}
          {/*  <Container style={{...homeStyles.infoContainer}}>*/}
          {/*    /!*@ts-expect-error - style not getting detected but gets picked up..*!/*/}
          {/*    <InfoIcon style={{color: "#54aeff"}}/>*/}
          {/*    <Text>Downloads are currently limited to 20 songs per day per user :)</Text>*/}
          {/*  </Container>*/}
          {/*</Container>*/}

          <TextInput
            style={{ marginTop: '3%' }}
            placeholder="Enter your URL here"
            value={url || ""}
            onChange={(e) => {
              setUrl(e.target.value)
            }}
          />

          <Container style={{ display: "flex", marginTop: "1%", alignItems: "center", marginLeft: "0", padding: 0, gap: "5%",  }}>
            <Button
            variant="light"
            style={{minWidth: "8rem"}}
            onClick={() => download(url ?? "")}
            >
              Download
            </Button>
          </Container>
          {downloading &&
            <Container style={{
              display: "flex",
              marginLeft: 0,
              marginTop: "1.5%",
              flexDirection: "column",
              alignItems: "flex-start",
              padding: 0,
              justifyItems: "center"
            }}>
              <Loader size={20}/>
              <Text>  Your download is being processed, please wait.</Text>
            </Container>
          }
          {downloadedPath &&
            <Container style={{
              display: "flex",
              marginLeft: 0,
              marginTop: "1.5%",
              flexDirection: "column",
              alignItems: "flex-start",
              padding: 0,
              justifyItems: "center"
            }}>
              <Text>Your download was successful! File saved to:<br/>{downloadedPath}</Text>
            </Container>
          }
        </Container>
      </AppShell.Main>
      <Footer/>
    </>
  )
}