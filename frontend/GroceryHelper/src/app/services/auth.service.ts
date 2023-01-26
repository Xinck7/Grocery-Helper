import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';
import { LoggedInUser } from './auth'

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  baseUrl = environment.base_url
 
  login(data: any): Observable<LoggedInUser> {
    return this.http.post(`${this.baseUrl}/api-user-login/`, data)
      .pipe(map(result => {
        localStorage.setItem('authUser', JSON.stringify(result));
        return result;
      })) as Observable<LoggedInUser>; ;
  }     

  register(data: any) {
    return this.http.post(`${this.baseUrl}/api/v1/users/`, data);
  }

  getProfile(userId: string|null): Observable<any> {
    return this.http.get(`${this.baseUrl}/api/v1/users/${userId}/`);
  }

  logout() {
    return this.http.get(`${this.baseUrl}/api/v1/logout/`)
      .pipe(tap(() => {
        localStorage.removeItem('authUser')
      }));
  }

  getAuthUser() {
    return JSON.parse(localStorage.getItem('authUser') as string);
  }
  
  setLoggedInUser(userData: LoggedInUser): void {
    if (localStorage.getItem('userData') !== JSON.stringify(userData)) {
      localStorage.setItem('userData', JSON.stringify(userData));
    }
   }

  get isLoggedIn() {
    if (localStorage.getItem('authUser')) {
      return true;
    }
    return false;
  }
}
