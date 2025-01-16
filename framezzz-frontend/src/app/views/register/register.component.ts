import { Component } from '@angular/core';
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

@Component({
  selector: 'app-register',
  imports: [
    ToolbarComponent,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.sass',
})
export class RegisterComponent {
  protected registerForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required]),
    repeatPassword: new FormControl('', [Validators.required]),
  });

  protected submit() {
    if (this.registerForm.invalid) return;

    const email = this.registerForm.get('email')?.value;
    const password = this.registerForm.get('password')?.value;
    const repeatPassword = this.registerForm.get('repeatPassword')?.value;

    if (password !== repeatPassword) return;
  }
}
