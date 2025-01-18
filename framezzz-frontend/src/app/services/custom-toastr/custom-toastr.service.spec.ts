import { TestBed } from '@angular/core/testing';
import { ToastrService, ToastPackage } from 'ngx-toastr';
import { CustomToastrService } from './custom-toastr.service';

describe('CustomToastrService', () => {
  let service: CustomToastrService;
  let toastrServiceMock: jasmine.SpyObj<ToastrService>;

  beforeEach(() => {
    toastrServiceMock = jasmine.createSpyObj('ToastrService', [
      'success',
      'error',
      'warning',
      'info',
    ]);

    TestBed.configureTestingModule({
      providers: [
        CustomToastrService,
        { provide: ToastrService, useValue: toastrServiceMock },
      ],
    });

    service = TestBed.inject(CustomToastrService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should call success method on ToastrService', () => {
    service.success('Test message', 'Test title');
    expect(toastrServiceMock.success).toHaveBeenCalledWith(
      'Test message',
      'Test title',
      jasmine.any(Object)
    );
  });

  it('should call error method on ToastrService', () => {
    service.error('Error message', 'Error title');
    expect(toastrServiceMock.error).toHaveBeenCalledWith(
      'Error message',
      'Error title',
      jasmine.any(Object)
    );
  });

  it('should call warning method on ToastrService', () => {
    service.warning('Warning message', 'Warning title');
    expect(toastrServiceMock.warning).toHaveBeenCalledWith(
      'Warning message',
      'Warning title',
      jasmine.any(Object)
    );
  });

  it('should call info method on ToastrService', () => {
    service.info('Info message', 'Info title');
    expect(toastrServiceMock.info).toHaveBeenCalledWith(
      'Info message',
      'Info title',
      jasmine.any(Object)
    );
  });
});
