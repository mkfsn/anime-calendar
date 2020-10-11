import { writable, Writable } from 'svelte/store';
import type { IAnime } from '../anime';

interface ISelectableAnime extends IAnime {
    selected: number[];
}

interface ISelectableAnimes {
    [key: string]: ISelectableAnime,
}

function createTimetables() {
    const { subscribe, set, update } = writable({} as ISelectableAnimes);
    return {
        subscribe,
        set,
        update,
        addProgram: (anime: IAnime, channelIndex: number) => update((o: ISelectableAnimes) => {
            if (!o[anime.Id]) {
                o[anime.Id] = {
                    ...anime,
                    selected: [],
                }
            }
            o[anime.Id].selected.push(channelIndex);
            return o;
        }),
        deleteProgram: (anime: IAnime, channelIndex: number) => update((o: ISelectableAnimes) => {
            o[anime.Id].selected = o[anime.Id].selected.filter(v => v !== channelIndex)
            if (o[anime.Id].selected.length === 0) {
                delete o[anime.Id];
            }
            return o;
        })
    };
}

const timetablesStore:Writable<ISelectableAnimes> = createTimetables();

export default timetablesStore;