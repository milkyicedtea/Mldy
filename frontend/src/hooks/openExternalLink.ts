import {BrowserOpenURL} from "@wailsjs/runtime";

export function openExternalLink(url: string) {
  BrowserOpenURL(url);
}