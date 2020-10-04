import {compressToEncodedURIComponent, decompressFromEncodedURIComponent} from 'lz-string';

export async function fetchAnime() {
    const response = await fetch(`/data/anime.json`);
    if (response.ok) {
        return response.json();
    }
    throw new Error("failed to fetch anime");
}

export function deserializeConfig(serializedConfig) {
    return JSON.parse(decompressFromEncodedURIComponent(serializedConfig))
}

export function serializeConfig(config) {
    return compressToEncodedURIComponent(JSON.stringify(config))
}