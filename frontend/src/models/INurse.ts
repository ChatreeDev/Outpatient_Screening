import { OutpatientScreeningsInterface } from "./IOutpatientScreening";

export interface NurseInterface {
    id: number;
    firstName: string;
    Email: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}