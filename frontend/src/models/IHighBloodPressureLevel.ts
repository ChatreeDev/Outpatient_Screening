import { OutpatientScreeningsInterface } from "./OutpatientScreening";

export interface HighBloodPressureLevelsInterface {
    ID: number;
    Name: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}