import { writable } from 'svelte/store';

function createTimetables() {
    const { subscribe, set, update } = writable({
        // [program]: {
        //  [channel]: []
        //}
    });
    return {
        subscribe,
        addProgram: (program, channel, timetables) => update(o => {
            if (!o[program]) {
                o[program] = {};
            }
            o[program][channel] = timetables;
            return o;
        }),
        deleteProgram: (program, channel) => update(o => {
            delete o[program][channel];
            if (o[program].length === 0) {
                delete o[program];
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