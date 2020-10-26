import fs from "fs";
import parser from "iptv-playlist-parser";

const playlist = fs.readFileSync("./test.m3u8", { encoding: "utf-8" });
const result = parser.parse(playlist);

console.log(result["items"][0]["tvg"]);
