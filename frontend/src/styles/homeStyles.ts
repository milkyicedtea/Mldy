import {CSSProperties} from "@mantine/core"

export const homeStyles: {[key: string]: CSSProperties} = {
  titleContainer: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  anchor: {
    textDecoration: "none"
  },
  title: {
    fontSize: "2rem",
    cursor: "pointer"
  },
  contentContainer: {
    marginTop: "1rem",
    display: "flex",
    flexDirection: "column",
    width: '100%',
    height: '80%'
  },
  infoGroupContainer: {
    margin: 0,
    display: "flex",
    padding: "16px",
    backgroundColor: "rgba(84,174,255,0.1)",
    borderLeft: "4px solid #54aeff",
    marginBottom: '.5rem',
    placeSelf: 'start',
    marginTop: '1%'
  },
  infoContainer: {
    padding: 0,
    display: "flex",
    gap: ".5rem",
    alignItems: "center"
  },
  iconContainer: {
    aspectRatio: 1,
    backgroundColor: '#fff', // or any background color to avoid transparent SVG issues
    borderRadius: '50%', // makes the icon circular
    width: '3.5rem',
    height: '3.5rem',
    padding: '10px',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.7)', // optional shadow for depth
    transition: 'all 0.3s ease',
  },
  platformIcon: {
    width: "2.5rem",
    height: "2.5rem",
    aspectRatio :1,
    objectFit: 'contain',
  },
  footerContainer: {
    display: "flex",
    flexDirection: "row",
    gap: "4%",
    alignItems: "center"
  }
}