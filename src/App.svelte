<script>
  import AnimeList from "./components/AnimeList.svelte";
  import Navbar from "./components/Navbar.svelte";
</script>

<style>
  main {
    align-content: center;
    box-sizing: border-box;
    font-family: 'SF Pro Text', 'SF Pro Icons', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;
    hyphens: auto;
    line-height: 1.65;
    margin: 0 auto;
    min-height: calc(100vh - 4em);
    text-align: center;
  }
</style>

<svelte:head>
  <title>Anime Calendar</title>
</svelte:head>

<Navbar title="アニメ番組"/>
<main>
  <AnimeList />
</main>
