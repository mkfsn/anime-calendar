<script>
  import AnimeList from "./components/AnimeList.svelte";
  import Navbar from "./components/Navbar.svelte";
  import Calendar from "./components/Calendar.svelte";

  let showAnimeList = true;
</script>

<style>
  main {
    align-content: center;
    box-sizing: border-box;
    font-family: 'SF Pro Text', 'SF Pro Icons', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;
    hyphens: auto;
    margin: 0 auto;
    text-align: center;
  }
</style>

<svelte:head>
  <title>Anime Calendar</title>
</svelte:head>

<Navbar title="アニメ番組" bind:active={showAnimeList}/>
<main>
  <AnimeList show={showAnimeList} />
  <Calendar />
</main>
