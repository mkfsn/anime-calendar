<script>
  import Anime from "./components/Anime.svelte";
  async function fetchAnime() {
      const response = await fetch(`/data/anime.json`);
      if (response.ok) {
        return response.json();
      }
      throw new Error("failed to fetch anime");
  }
  let promise = fetchAnime();
</script>

<style>
  .scrollable {
    width: 100%;
    overflow-x: scroll;
    background-color: #e2e2e2;
    padding-left: 1em;
    padding-right: 1em;
  }
  .container {
    display: flex;
    flex-direction: row;
    float: left;
  }
</style>

<svelte:head>
  <title>Anime Calendar</title>
</svelte:head>

<main>
  <h1>Anime</h1>
  {#await promise}
    <p>Loading ...</p>
  {:then animes}
    <div class="scrollable">
      <div class="container">
        {#each animes as anime, i}
          <Anime {anime} />
        {/each}
      </div>
    </div>
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</main>
