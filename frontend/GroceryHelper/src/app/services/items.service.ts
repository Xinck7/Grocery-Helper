import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';


@Injectable({
  providedIn: 'root'
})
export class ItemsService {

  constructor(private http: HttpClient) { }

  baseUrl = environment.base_url

  getItems() {
    return this.http.get(`${this.baseUrl}/api/v1/items/`);
  }

  // todo change/add item model to get rid of any type
  createItem(data: any){
    return this.http.post(`${this.baseUrl}/api/v1/items/`, data);
  }

}
