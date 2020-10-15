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
    EntPositionassingment,
    EntPositionassingmentFromJSON,
    EntPositionassingmentFromJSONTyped,
    EntPositionassingmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPhysicianEdges
 */
export interface EntPhysicianEdges {
    /**
     * Formuser holds the value of the formuser edge.
     * @type {Array<EntPositionassingment>}
     * @memberof EntPhysicianEdges
     */
    formuser?: Array<EntPositionassingment>;
}

export function EntPhysicianEdgesFromJSON(json: any): EntPhysicianEdges {
    return EntPhysicianEdgesFromJSONTyped(json, false);
}

export function EntPhysicianEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPhysicianEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'formuser': !exists(json, 'formuser') ? undefined : ((json['formuser'] as Array<any>).map(EntPositionassingmentFromJSON)),
    };
}

export function EntPhysicianEdgesToJSON(value?: EntPhysicianEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'formuser': value.formuser === undefined ? undefined : ((value.formuser as Array<any>).map(EntPositionassingmentToJSON)),
    };
}


