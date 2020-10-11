import {Writable, writable} from "svelte/store";

interface ISelectedPrograms {
    [key: number]: number[];
}

interface IConfigStore {
    programs: ISelectedPrograms;
}

class ConfigStore implements IConfigStore {
    programs: ISelectedPrograms;
    constructor() {
        this.programs = [];
    }
}

function createConfigStore() {
    const { subscribe, set, update } = writable(new ConfigStore);
    return {
        subscribe,
        set,
        update,
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

const configStore:Writable<IConfigStore> = createConfigStore();

export default configStore;
