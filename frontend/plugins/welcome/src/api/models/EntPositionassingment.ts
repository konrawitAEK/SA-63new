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

import { exists, mapValues } from '../runtime';
import {
    EntPositionassingmentEdges,
    EntPositionassingmentEdgesFromJSON,
    EntPositionassingmentEdgesFromJSONTyped,
    EntPositionassingmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPositionassingment
 */
export interface EntPositionassingment {
    /**
     * DayStart holds the value of the "DayStart" field.
     * @type {string}
     * @memberof EntPositionassingment
     */
    dayStart?: string;
    /**
     * 
     * @type {number}
     * @memberof EntPositionassingment
     */
    departmentID?: number;
    /**
     * 
     * @type {EntPositionassingmentEdges}
     * @memberof EntPositionassingment
     */
    edges?: EntPositionassingmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPositionassingment
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntPositionassingment
     */
    physicianID?: number;
    /**
     * 
     * @type {number}
     * @memberof EntPositionassingment
     */
    positionID?: number;
}

export function EntPositionassingmentFromJSON(json: any): EntPositionassingment {
    return EntPositionassingmentFromJSONTyped(json, false);
}

export function EntPositionassingmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPositionassingment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dayStart': !exists(json, 'DayStart') ? undefined : json['DayStart'],
        'departmentID': !exists(json, 'department_ID') ? undefined : json['department_ID'],
        'edges': !exists(json, 'edges') ? undefined : EntPositionassingmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'physicianID': !exists(json, 'physician_ID') ? undefined : json['physician_ID'],
        'positionID': !exists(json, 'position_ID') ? undefined : json['position_ID'],
    };
}

export function EntPositionassingmentToJSON(value?: EntPositionassingment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'DayStart': value.dayStart,
        'department_ID': value.departmentID,
        'edges': EntPositionassingmentEdgesToJSON(value.edges),
        'id': value.id,
        'physician_ID': value.physicianID,
        'position_ID': value.positionID,
    };
}


