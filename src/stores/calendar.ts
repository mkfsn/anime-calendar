import {writable, Writable} from "svelte/store";

const calendarStore:Writable<Date> = writable(new Date());

export default calendarStore;