export interface Application {
    id: number;
    company: string;
    position: string;
    status: ApplicationStatus;
    location: string;
    applied_date: string;
    term: string;
    note?: string;
    resume_url: string;
    user_id: number;
}

export enum ApplicationStatus {
    Applied = 0,
    OAReceived = 1,
    Interviewing = 2,
    Accepted = 3,
    Rejected = 4
}

export interface LoadingState {
    isLoading: boolean;
    error: string | null;
}

export interface ApiResponse<T = any> {
    data?: T;
    error?: string;
    success?: boolean;
}

export const STATUS_COLORS = {
    [ApplicationStatus.Applied]: 'primary',
    [ApplicationStatus.OAReceived]: 'info',
    [ApplicationStatus.Interviewing]: 'warning',
    [ApplicationStatus.Accepted]: 'success',
    [ApplicationStatus.Rejected]: 'danger'
};
