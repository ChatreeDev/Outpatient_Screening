import { OutpatientScreeningsInterface } from "./IOutpatientScreening";

export interface ObesityLevelInterface {
    ID: number;
    Level: string;
    AssessmentForms: string;
    HistoryTakingForms: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}