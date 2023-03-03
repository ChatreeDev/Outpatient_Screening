import { OutpatientScreeningsInterface } from "./IOutpatientScreenings";

export interface EmployeesInterface {
    id: number;
    firstName: string;
    Email: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}