import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ListsService {

  constructor(private http: HttpClient) { }

  baseUrl = environment.base_url
  
  getLists() {
    return this.http.get(`${this.baseUrl}/api/v1/list/`);
  }
  
}
