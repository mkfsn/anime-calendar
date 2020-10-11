<script>
    import Navbar from "./components/Navbar.svelte";
    import CalendarView from "./components/Calendar/CalendarView.svelte";
    import Slider from "./components/Slider.svelte";
    import AnimeList from "./components/Anime/AnimeList.svelte";
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
    <Slider collapse={!showAnimeList}>
        <AnimeList />
    </Slider>
    <CalendarView />
</main>
