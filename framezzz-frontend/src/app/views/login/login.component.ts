import { Component, inject } from '@angular/core';
import { ToolbarComponent } from '../../components/toolbar/toolbar.component';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AuthService } from '../../services/auth/auth.service';
import { CookieService } from 'ngx-cookie-service';
import { Router } from '@angular/router';
import { UserData } from '../../interfaces/user-data.interface';
import { LoginResponse } from '../../interfaces/login-response.interface';
import { CustomToastrService } from '../../services/custom-toastr/custom-toastr.service';

@Component({
  selector: 'app-login',
  imports: [
    ToolbarComponent,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.sass',
})
export class LoginComponent {
  private readonly router = inject(Router);

  private readonly cookieService = inject(CookieService);

  private readonly authService = inject(AuthService);

  private customToastrService = inject(CustomToastrService);

  protected loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),

    password: new FormControl('', [Validators.required]),
  });

  protected submit() {
    if (this.loginForm.invalid) return;

    const email = this.loginForm.get('email')?.value;

    const password = this.loginForm.get('password')?.value;

    const user: UserData = { email: email!, password: password! };

    this.authService.loginUser(user).subscribe({
      next: (response: LoginResponse) => {
        this.cookieService.set('authToken', response.token, 7);

        this.router.navigate(['']);
      },
      error: () => {
        this.customToastrService.error('Failed to login user.');
      },
    });
  }
}
