import { OutpatientScreeningsInterface } from "./IOutpatientScreening";

export interface HighBloodPressureLevelsInterface {
    ID: number;
    Level: string;
    AssessmentForms: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}