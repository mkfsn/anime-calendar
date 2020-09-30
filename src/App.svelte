<script>
  import AnimeList from "./components/AnimeList.svelte";
</script>

<style>
  main {
    align-content: center;
    box-sizing: border-box;
    font-family: 'SF Pro Text', 'SF Pro Icons', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;
    hyphens: auto;
    line-height: 1.65;
    margin: 0 auto;
    max-width: 1200px;
    min-height: 100vh;
    text-align: center;
  }
</style>

<svelte:head>
  <title>Anime Calendar</title>
</svelte:head>

<main>
  <h1>Anime</h1>
  <AnimeList />
</main>
