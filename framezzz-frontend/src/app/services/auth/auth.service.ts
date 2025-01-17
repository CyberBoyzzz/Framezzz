import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly httpClient = inject(HttpClient);

  public registerUser(email: string, password: string) {
    return this.httpClient.post(`/api/auth/register`, { email, password });
  }

  public loginUser(email: string, password: string): Observable<any> {
    return this.httpClient.post(`/api/auth/login`, { email, password });
  }
}
