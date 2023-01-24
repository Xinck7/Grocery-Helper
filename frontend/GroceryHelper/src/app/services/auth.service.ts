import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

   baseUrl = environment.base_url
 
  login(data: any) {
    return this.http.post(`${this.baseUrl}/login`, data)
      .pipe(map(result => {
        localStorage.setItem('authUser', JSON.stringify(result));
        return result;
      }));
  }
  
  register(data: any) {
    return this.http.post(`${this.baseUrl}/api/v1/users/`, data);
   }
 
  profile(userId: string|null): Observable<any> {
    return this.http.get(`${this.baseUrl}/api/v1/users/${userId}/`);
  }

  logout() {
    return this.http.get(`${this.baseUrl}/api/v1/logout`)
      .pipe(tap(() => {
        localStorage.removeItem('authUser')
      }));
  }

  getAuthUser() {
    return JSON.parse(localStorage.getItem('authUser') as string);
  }

  get isLoggedIn() {
    if (localStorage.getItem('authUser')) {
      return true;
    }
    return false;
  }
}
