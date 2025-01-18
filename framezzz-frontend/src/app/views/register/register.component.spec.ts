import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RegisterComponent } from './register.component';
import { Router } from '@angular/router';
import { CustomToastrService } from '../../services/custom-toastr/custom-toastr.service';
import { ToastrService } from 'ngx-toastr';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;

  let toastrServiceMock: jasmine.SpyObj<CustomToastrService>;
  let routerMock: jasmine.SpyObj<Router>;

  beforeEach(async () => {
    toastrServiceMock = jasmine.createSpyObj('CustomToastrService', [
      'success',
      'error',
    ]);
    routerMock = jasmine.createSpyObj('Router', ['navigate']);

    await TestBed.configureTestingModule({
      imports: [
        RegisterComponent,
        HttpClientTestingModule,
        BrowserAnimationsModule,
      ],
      providers: [
        { provide: CustomToastrService, useValue: toastrServiceMock },
        { provide: Router, useValue: routerMock },
        ToastrService,
        {
          provide: 'ToastConfig',
          useValue: {
            timeOut: 3000,
            toastClass: 'toast-new',
            positionClass: 'toast-bottom-right',
            progressBar: true,
          },
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
