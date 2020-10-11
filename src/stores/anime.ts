import {writable} from "svelte/store";

const animeStore = writable({
    loaded: false,
    data: [],
});

export default animeStore;