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
    Interview = 1,
    OfferReceived = 2,
    Rejected = 3
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
    [ApplicationStatus.OfferReceived]: 'success',
    [ApplicationStatus.Rejected]: 'danger'
};
