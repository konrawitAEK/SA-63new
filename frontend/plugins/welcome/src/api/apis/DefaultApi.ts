/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    ControllersPositionassingment,
    ControllersPositionassingmentFromJSON,
    ControllersPositionassingmentToJSON,
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentToJSON,
    EntPhysician,
    EntPhysicianFromJSON,
    EntPhysicianToJSON,
    EntPosition,
    EntPositionFromJSON,
    EntPositionToJSON,
    EntPositionassingment,
    EntPositionassingmentFromJSON,
    EntPositionassingmentToJSON,
} from '../models';

export interface CreateDepartmentRequest {
    department: EntDepartment;
}

export interface CreatePhysicianRequest {
    physician: EntPhysician;
}

export interface CreatePositionRequest {
    position: EntPosition;
}

export interface CreatePositionassingmentRequest {
    positionassingment: ControllersPositionassingment;
}

export interface DeletePositionassingmentRequest {
    id: number;
}

export interface GetDepartmentRequest {
    id: number;
}

export interface GetPhysicianRequest {
    id: number;
}

export interface GetPositionRequest {
    id: number;
}

export interface GetPositionassingmentRequest {
    id: number;
}

export interface ListDepartmentRequest {
    limit?: number;
    offset?: number;
}

export interface ListPhysicianRequest {
    limit?: number;
    offset?: number;
}

export interface ListPositionRequest {
    limit?: number;
    offset?: number;
}

export interface ListPositionassingmentRequest {
    limit?: number;
    offset?: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create department
     * Create department
     */
    async createDepartmentRaw(requestParameters: CreateDepartmentRequest): Promise<runtime.ApiResponse<EntDepartment>> {
        if (requestParameters.department === null || requestParameters.department === undefined) {
            throw new runtime.RequiredError('department','Required parameter requestParameters.department was null or undefined when calling createDepartment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/departments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntDepartmentToJSON(requestParameters.department),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDepartmentFromJSON(jsonValue));
    }

    /**
     * Create department
     * Create department
     */
    async createDepartment(requestParameters: CreateDepartmentRequest): Promise<EntDepartment> {
        const response = await this.createDepartmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create physician
     * Create physician
     */
    async createPhysicianRaw(requestParameters: CreatePhysicianRequest): Promise<runtime.ApiResponse<EntPhysician>> {
        if (requestParameters.physician === null || requestParameters.physician === undefined) {
            throw new runtime.RequiredError('physician','Required parameter requestParameters.physician was null or undefined when calling createPhysician.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/physicians`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPhysicianToJSON(requestParameters.physician),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPhysicianFromJSON(jsonValue));
    }

    /**
     * Create physician
     * Create physician
     */
    async createPhysician(requestParameters: CreatePhysicianRequest): Promise<EntPhysician> {
        const response = await this.createPhysicianRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create position
     * Create position
     */
    async createPositionRaw(requestParameters: CreatePositionRequest): Promise<runtime.ApiResponse<EntPosition>> {
        if (requestParameters.position === null || requestParameters.position === undefined) {
            throw new runtime.RequiredError('position','Required parameter requestParameters.position was null or undefined when calling createPosition.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/positions`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPositionToJSON(requestParameters.position),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPositionFromJSON(jsonValue));
    }

    /**
     * Create position
     * Create position
     */
    async createPosition(requestParameters: CreatePositionRequest): Promise<EntPosition> {
        const response = await this.createPositionRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create positionassingment
     * Create positionassingment
     */
    async createPositionassingmentRaw(requestParameters: CreatePositionassingmentRequest): Promise<runtime.ApiResponse<ControllersPositionassingment>> {
        if (requestParameters.positionassingment === null || requestParameters.positionassingment === undefined) {
            throw new runtime.RequiredError('positionassingment','Required parameter requestParameters.positionassingment was null or undefined when calling createPositionassingment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/positionassingments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersPositionassingmentToJSON(requestParameters.positionassingment),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ControllersPositionassingmentFromJSON(jsonValue));
    }

    /**
     * Create positionassingment
     * Create positionassingment
     */
    async createPositionassingment(requestParameters: CreatePositionassingmentRequest): Promise<ControllersPositionassingment> {
        const response = await this.createPositionassingmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get positionassingment by ID
     * Delete a positionassingment entity by ID
     */
    async deletePositionassingmentRaw(requestParameters: DeletePositionassingmentRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deletePositionassingment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/positionassingments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get positionassingment by ID
     * Delete a positionassingment entity by ID
     */
    async deletePositionassingment(requestParameters: DeletePositionassingmentRequest): Promise<object> {
        const response = await this.deletePositionassingmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get department by ID
     * Get a department entity by ID
     */
    async getDepartmentRaw(requestParameters: GetDepartmentRequest): Promise<runtime.ApiResponse<EntDepartment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getDepartment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/departments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDepartmentFromJSON(jsonValue));
    }

    /**
     * get department by ID
     * Get a department entity by ID
     */
    async getDepartment(requestParameters: GetDepartmentRequest): Promise<EntDepartment> {
        const response = await this.getDepartmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get physician by ID
     * Get a physician entity by ID
     */
    async getPhysicianRaw(requestParameters: GetPhysicianRequest): Promise<runtime.ApiResponse<EntPhysician>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPhysician.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/physicians/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPhysicianFromJSON(jsonValue));
    }

    /**
     * get physician by ID
     * Get a physician entity by ID
     */
    async getPhysician(requestParameters: GetPhysicianRequest): Promise<EntPhysician> {
        const response = await this.getPhysicianRaw(requestParameters);
        return await response.value();
    }

    /**
     * get position by ID
     * Get a position entity by ID
     */
    async getPositionRaw(requestParameters: GetPositionRequest): Promise<runtime.ApiResponse<EntPosition>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPosition.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/positions/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPositionFromJSON(jsonValue));
    }

    /**
     * get position by ID
     * Get a position entity by ID
     */
    async getPosition(requestParameters: GetPositionRequest): Promise<EntPosition> {
        const response = await this.getPositionRaw(requestParameters);
        return await response.value();
    }

    /**
     * get positionassingment by ID
     * Get a positionassingment entity by ID
     */
    async getPositionassingmentRaw(requestParameters: GetPositionassingmentRequest): Promise<runtime.ApiResponse<EntPositionassingment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPositionassingment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/positionassingments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPositionassingmentFromJSON(jsonValue));
    }

    /**
     * get positionassingment by ID
     * Get a positionassingment entity by ID
     */
    async getPositionassingment(requestParameters: GetPositionassingmentRequest): Promise<EntPositionassingment> {
        const response = await this.getPositionassingmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * list department entities
     * List department entities
     */
    async listDepartmentRaw(requestParameters: ListDepartmentRequest): Promise<runtime.ApiResponse<Array<EntDepartment>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/departments`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntDepartmentFromJSON));
    }

    /**
     * list department entities
     * List department entities
     */
    async listDepartment(requestParameters: ListDepartmentRequest): Promise<Array<EntDepartment>> {
        const response = await this.listDepartmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * list physician entities
     * List physician entities
     */
    async listPhysicianRaw(requestParameters: ListPhysicianRequest): Promise<runtime.ApiResponse<Array<EntPhysician>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/physicians`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPhysicianFromJSON));
    }

    /**
     * list physician entities
     * List physician entities
     */
    async listPhysician(requestParameters: ListPhysicianRequest): Promise<Array<EntPhysician>> {
        const response = await this.listPhysicianRaw(requestParameters);
        return await response.value();
    }

    /**
     * list position entities
     * List position entities
     */
    async listPositionRaw(requestParameters: ListPositionRequest): Promise<runtime.ApiResponse<Array<EntPosition>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/positions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPositionFromJSON));
    }

    /**
     * list position entities
     * List position entities
     */
    async listPosition(requestParameters: ListPositionRequest): Promise<Array<EntPosition>> {
        const response = await this.listPositionRaw(requestParameters);
        return await response.value();
    }

    /**
     * list positionassingment entities
     * List positionassingment entities
     */
    async listPositionassingmentRaw(requestParameters: ListPositionassingmentRequest): Promise<runtime.ApiResponse<Array<EntPositionassingment>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/positionassingments`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPositionassingmentFromJSON));
    }

    /**
     * list positionassingment entities
     * List positionassingment entities
     */
    async listPositionassingment(requestParameters: ListPositionassingmentRequest): Promise<Array<EntPositionassingment>> {
        const response = await this.listPositionassingmentRaw(requestParameters);
        return await response.value();
    }

}
