interface IAnimeSong {
   Title: string;
   Details: { [key: string]: string };
}

interface IProgram {
   Time: string;
   Name: string;
}

interface ILink {
   Title: string;
   URL: string;
}

interface ITimetable {
   Channel: string;
   StartAt: string;
   Duration: number;
   Number: number;
   Subtitle: string;
}

interface IAnime {
   Id: string;
   Title: string;
   Website: string;
   OpenGraph?: any;
   Staff: { [key: string]: string };
   Opening?: IAnimeSong;
   Ending?: IAnimeSong;
   Cast: { [key: string]: string };
   Programs: IProgram[];
   Links: ILink[];
   Channels: string[];
   Timetables: ITimetable[];
}

export type { IAnime, ITimetable };