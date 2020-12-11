// routing.js
import { ApiURl } from './utils'; 
const APIURL = ApiURl();

export const AidataListUrl = APIURL + '/list';              // get everything 
export const AidataTotalCountUrl = APIURL + '/count';       // get total row count
export const AidataFilteredListUrl = APIURL + '/eyes/';     // append :langcode/:monthYear/:recorder
export const AidataByKeycodeUrl = APIURL + '/key/';         // append keycode
export const AidataUpdateUrl = APIURL + '/key/';            // append keycode/{up|down}
export const RecorderCodesUrl = APIURL + '/codes';          // get all recorderCodes

// Served as static files in routing.go :
export const PrivacyPolicyUrl = 'PrivacyPolicy.html';
