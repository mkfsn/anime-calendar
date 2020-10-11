import { writable } from 'svelte/store';

function createTimetables() {
    const { subscribe, set, update } = writable({
        // [id]: {
        //   ...anime,
        //   [selected]: []
        // }
    });
    return {
        subscribe,
        addProgram: (anime, channelIndex) => update(o => {
            if (!o[anime.Id]) {
                o[anime.Id] = {
                    ...anime,
                    ["selected"]: [],
                }
            }
            o[anime.Id]["selected"].push(channelIndex);
            return o;
        }),
        deleteProgram: (anime, channelIndex) => update(o => {
            o[anime.Id]["selected"] = o[anime.Id]["selected"].filter(v => v !== channelIndex)
            if (o[anime.Id]["selected"].length === 0) {
                delete o[anime.Id];
            }
            return o;
        })
    };
}

function createConfigStore() {
    const { subscribe, set, update } = writable({
        "programs": {
            // [program]: [...channel]
        }
    });
    return {
        subscribe,
        set,
        addProgram: (program, channel) => update(o => {
            if (!o.programs[program]) {
                o.programs[program] = [];
            }
            o.programs[program] = o.programs[program].filter(c => c !== channel)
            o.programs[program].push(channel)
            return o;
        }),
        deleteProgram: (program, channel) => update(o => {
            o.programs[program] = o.programs[program].filter(c => c !== channel)
            if (o.programs[program].length === 0) {
                delete o.programs[program];
            }
            return o;
        })
    };
}

export const configStore = createConfigStore();
export const dateStore = writable(new Date());
export const timetables = createTimetables();
export const animeStore = writable({
    loaded: false,
    data: [],
});