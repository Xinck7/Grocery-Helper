import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class IngredientsService {

  constructor(private http: HttpClient) { }

  baseUrl = environment.base_url

  getIngredients() {
    return this.http.get(`${this.baseUrl}/api/v1/ingredients/`);
  }

}
