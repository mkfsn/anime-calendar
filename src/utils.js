export async function fetchAnime() {
    const response = await fetch(`/data/anime.json`);
    if (response.ok) {
        return response.json();
    }
    throw new Error("failed to fetch anime");
}